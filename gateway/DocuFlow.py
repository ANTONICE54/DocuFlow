import openai
from gateway.config import *
import docx
from fastapi import FastAPI, UploadFile, File, HTTPException
from pydantic import BaseModel
from pypdf import PdfReader



api_app = FastAPI()

class extract_data_from_file_result(BaseModel):
    document_text: str

class get_explanation_result(BaseModel):
    explenation_result: str


@api_app.post("/extract_data_from_file/")
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



@api_app.get("/get_explenation")
async def get_explenation(document_text: str) -> get_explanation_result:
    try:
        chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)

        explenation_result = chat_gpt_request.chat.completions.create(
            model="gpt-4o",
            messages=[
                {"role": "system", "content":document_text},
                {"role": "user", "content": '(Country Ukraine) Explain this document in simple words for an ordinary person, but do not forget about the details (names, numbers, etc.). As a result, you will receive a text about this document in the language of the country.'}
            ]
        ).choices[0].message.content
    except:
        raise HTTPException(400, detail="Explanation was failed")

    return {"explenation_result": explenation_result}

# document_text =  '\n'.join(paragraph.text for paragraph in docx.Document('/Users/dmitro/Downloads/Договір з Ліцензіатом.docx').paragraphs)

# chat_gpt_request = openai.OpenAI(api_key=CHAT_GPT_TOKEN)


# completion = chat_gpt_request.chat.completions.create(
#     model="gpt-4o",
#     messages=[
#         {"role": "system", "content":document_text},
#         {"role": "assistant", "content": '(Country Ukraine) Explain this document in simple words for an ordinary person, but do not forget about the details (names, numbers, etc.). As a result, you will receive a text about this document in the language of the country.'}
#     ]
# ).choices[0].message.content

# print(completion)