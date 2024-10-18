from main_functions import *


def get_categories_and_subcategories(user_id):
    categories = make_get_request(f"http://localhost:8081/category", json={"user_id": user_id})
    if not check_for_error(categories):categories = categories.get("category_list", {'error':''}) 
    
    return categories


def create_category_function(user_id, category_name):
    create_category_response = make_post_request(f"http://localhost:8081/category", json={"user_id": user_id, "name":category_name})
    
    return create_category_response


def create_subcategory_function(category_id, subcategory_name):
    create_subcategory_response = make_post_request(f"http://localhost:8081/subcategory", json={"category_id": category_id, "name":subcategory_name})
    
    return create_subcategory_response