from main_functions import *


def get_categories_and_subcategories(user_id):
    categories = make_get_request(f"http://localhost:8081/category", json={"user_id": user_id})
    if type(categories)!=str: categories = categories.get("category_list", 'error') 
    
    return categories

def create_category(user_id, category_name):
    create_category_response = make_post_request(f"http://localhost:8081/category", json={"user_id": user_id, "name":category_name})
    
    return create_category_response