from main_functions import *


def get_user_info(user_id):
    user_info = make_get_request(f"http://localhost:8080/user/{user_id}")
    
    return user_info