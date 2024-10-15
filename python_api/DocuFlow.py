import openai
from config import *
import docx
from fastapi import FastAPI, UploadFile, File, Depends, FastAPI, HTTPException, APIRouter
from typing import Optional
from fastapi.security.http import HTTPAuthorizationCredentials, HTTPBearer
from pydantic import BaseModel, UUID4
from pypdf import PdfReader

api_app = FastAPI()

auth_token_handler = HTTPBearer(auto_error=False)

async def get_token(auth: Optional[HTTPAuthorizationCredentials] = Depends(auth_token_handler)) -> str:
    if auth is None or auth.credentials != AUTH_SWAGGER_TOKEN:
        raise HTTPException(401, detail="Bearer token missing or unknown")
    return auth.credentials


docu_flow_router = APIRouter(
    tags=["DocuFlow"],
    dependencies=[Depends(get_token)]
)



class extract_data_from_file_result(BaseModel):
    document_text: str

class get_explanation_result(BaseModel):
    explanation_result: str

class replace_result(BaseModel):
    replace_result: str

class add_result(BaseModel):
    add_result: str

class create_result(BaseModel):
    create_result: str

class create_tags(BaseModel):
    tags: list[str]


@docu_flow_router.post("/extract_data_from_file/")
async def extract_data_from_file(input_file: UploadFile = File(...)) -> extract_data_from_file_result:
    if input_file.content_type == "application/pdf":
        document_text =  '\n'.join(pdf_page.extract_text() for pdf_page in PdfReader(input_file.file).pages)
    elif input_file.content_type == "text/plain":
        document_text = input_file.file.read().decode('utf-8')
    elif input_file.content_type == "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
        document_text =  '\n'.join(paragraph.text for paragraph in docx.Document(input_file.file).paragraphs)
    else:
        raise HTTPException(400, detail="Invalid document type")
    return {"document_text": document_text}



@docu_flow_router.post("/get_explanation")
async def get_explanation(document_text: UploadFile) -> get_explanation_result:
    if document_text.content_type != "text/plain":
        raise HTTPException(400, detail="document_text must be .txt")
    try:
        document_text = document_text.file.read().decode('utf-8')
        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        explanation_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'document:{document_text}'},
                {"role": "assistant", "content": f'(Country Ukraine).\nMake explanation of this document in simple words for an ordinary person by points, but do not forget about the details (names, numbers, etc.). As a result, you will receive a text about this document in the language of the country. Finally, make it look good. Make more than 2000 characters'}
            ]
        ).choices[0].message.content
    except Exception as explanation_error:
        raise HTTPException(400, detail=f"Explanation was failed\n\n{explanation_error}")

    return {"explanation_result": explanation_result}



@docu_flow_router.post("/get_explanation_with_request")
async def get_explanation_with_request(document_text: UploadFile, explain_request: str) -> get_explanation_result:
    if document_text.content_type != "text/plain":
        raise HTTPException(400, detail="document_text must be .txt")

    try:
        document_text = document_text.file.read().decode('utf-8')

        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        explanation_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'document:{document_text}'},
                {"role": "system", "content": f'explain request:{explain_request}'},
                {"role": "assistant", "content": f'(Country Ukraine).\nMake explanation of this documentusing explain request in simple words for an ordinary person by points, but do not forget about the details (names, numbers, etc.). As a result, you will receive a text about this document in the language of the country. Finally, make it look good. Make more than 800 characters'}
            ]
        ).choices[0].message.content
    except Exception as explanation_error:
        raise HTTPException(400, detail=f"Explanation was failed\n\n{explanation_error}")

    return {"explanation_result": explanation_result}



@docu_flow_router.post("/get_part_explanation")
async def get_part_explanation(document_text: UploadFile, explain_part: UploadFile) -> get_explanation_result:
    if document_text.content_type != "text/plain":
        raise HTTPException(400, detail="document_text must be .txt")
    elif explain_part.content_type != "text/plain":
        raise HTTPException(400, detail="explain_part must be .txt")

    try:
        document_text = document_text.file.read().decode('utf-8')
        explain_part = explain_part.file.read().decode('utf-8')
        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        explanation_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'document:{document_text}'},
                {"role": "system", "content": f'explain part:{explain_part}'},
                {"role": "assistant", "content": f'(Country Ukraine).\nMake very easy explanation of this document by user explanaition request in super simple words for an ordinary person, but do not forget about the details (names, numbers, etc.). As a result, you will receive a text about this document in the language of the country. Finally, make it look good. Make more than 500 characters'}
            ]
        ).choices[0].message.content
    except Exception as explanation_error:
        raise HTTPException(400, detail=f"Explanation was failed\n\n{explanation_error}")

    return {"explanation_result": explanation_result}



@docu_flow_router.post("/replace_document_part")
async def replace_document_part(document_text: UploadFile, replace_document_part: UploadFile, replace_request: str = "") -> replace_result:
    if document_text.content_type != "text/plain":
        raise HTTPException(400, detail="document_text must be .txt")
    elif replace_document_part.content_type != "text/plain":
        raise HTTPException(400, detail="replace_document_part must be .txt")

    try:
        document_text = document_text.file.read().decode('utf-8')
        replace_document_part = replace_document_part.file.read().decode('utf-8')

        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        replace_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'document:{document_text}'},
                {"role": "system", "content": f'replace request:{replace_request}'},
                {"role": "system", "content": f'Part for replace:{replace_document_part}'},
                {"role": "assistant", "content": f'(Country Ukraine).\nIf replacement request is exist, use it. Replace part to make it better for all sides good in the document context. Keep context but replace all this part. As a result, you will receive a replaced text in the language of the country. All data that you do not have, make it in the <[data]> format, make the name of this data clear. Finally, make it look good. Create roughly the same number of characters'}
            ]
        ).choices[0].message.content
    except Exception as replace_error:
        raise HTTPException(400, detail=f"Replacce was failed\n\n{replace_error}")

    return {"replace_result": replace_result}



@docu_flow_router.post("/add_document_part")
async def replace_document_part(document_text: UploadFile, add_request: str) -> add_result:
    if document_text.content_type != "text/plain":
        raise HTTPException(400, detail="document_text must be .txt")

    try:
        document_text = document_text.file.read().decode('utf-8')

        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        add_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'document:{document_text}'},
                {"role": "system", "content": f'add request:{add_request}'},
                {"role": "assistant", "content": f'(Country Ukraine).\Add part to document using add request to make it better for all sides good in the document context. As a result, you will receive only the text that I need add to document in the language of the country. All data that you do not have, make it in the <[data]> format, make the name of this data clear. Finally, make it look good. Do not make it small.'}
            ]
        ).choices[0].message.content
    except Exception as add_error:
        raise HTTPException(400, detail=f"Add was failed\n\n{add_error}")

    return {"add_result": add_result}



@docu_flow_router.post("/create_document_by_template")
async def create_document_by_template(template_document_text: UploadFile, create_request: str) -> create_result:
    if template_document_text.content_type != "text/plain":
        raise HTTPException(400, detail="template_document_text must be .txt")

    try:
        template_document_text = template_document_text.file.read().decode('utf-8')

        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        create_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'template document:{template_document_text}'},
                {"role": "system", "content": f'createrequest:{create_request}'},
                {"role": "assistant", "content": f'(Country Ukraine) Create a document that matches the request and will be good for all parties. As a result, indicate only the document in the language of the country without any tips. All data that you do not have, make it in the <[data]> format, make the name of this data clear. Finally, make it look good. Do not make it small. Create roughly the same number of characters, more than 6000 characters'}
            ]
        ).choices[0].message.content
    except Exception as create_error:
        raise HTTPException(400, detail=f"Create was failed\n\n{create_error}")

    return {"create_result": create_result}



@docu_flow_router.post("/create_tags")
async def create_tags_function(document_text: UploadFile) -> create_tags:
    if document_text.content_type != "text/plain":
        raise HTTPException(400, detail="document_text must be .txt")

    try:
        document_text = document_text.file.read().decode('utf-8')

        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        create_result = eval(chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content": f'document:{document_text}'},
                {"role": "assistant", "content": f'(Country Ukraine) Read the document. Add 10 tags that are most relevant to this document in country language. With this tag you can easy describe document. Do it in ["tag", "tag", "tag"] format because I can turn it into a list'}
            ]
        ).choices[0].message.content)
        print(create_result)
    except Exception as create_error:
        raise HTTPException(400, detail=f"Create was failed\n\n{create_error}")

    return {"tags": create_result}




api_app.include_router(docu_flow_router)