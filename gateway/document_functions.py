from main_functions import *
from config import Config





def get_document_text(document):
    document_text = make_post_request(f"http://localhost:5001/extract_data_from_file", files={"input_file": document}, headers = {'Authorization': f'Bearer {Config.PYTHON_API_TOKEN}'})
    if not check_for_error(document_text): document_text = document_text.get("document_text")
    
    return document_text


def save_document_template(user_id, document_name, document_text):
    create_folder(f'document_templates/{user_id}')

    document_name = get_unique_name(f'document_templates/{user_id}', document_name, '.txt')

    with open(f'document_templates/{user_id}/{document_name}.txt', "w") as document_template:
        document_template.write(document_text)

    return document_name


def save_document(user_id, document_name, document_text):
    create_folder(f'documents/{user_id}')

    document_name = get_unique_name(f'documents/{user_id}', document_name, '.txt')

    with open(f'documents/{user_id}/{document_name}.txt', "w") as document:
        document.write(document_text)

    return document_name


def create_document_by_template(user_id, template_document_name, create_request):
    create_result = make_post_request(f"http://localhost:5001/create_document_by_template", params={'create_request':create_request}, files={"template_document_text": (f"{template_document_name}.txt", open(f'document_templates/{user_id}/{template_document_name}.txt', 'rb'), 'text/plain')}, headers = {'Authorization': f'Bearer {Config.PYTHON_API_TOKEN}'})
    if not check_for_error(create_result): create_result = create_result.get("create_result")
    
    return create_result