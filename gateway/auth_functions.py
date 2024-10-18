from main_functions import *

def register_profile(name, surname, email, country, password):
    token = make_post_request("http://localhost:8080/register", json = {"name":name, "surname":surname, "email":email, "country":country, "password":password})
    if not check_for_error(token):token = token.get("token", {'error':''}) 

    return token

def log_in_profile(email, password):
    token = make_post_request("http://localhost:8080/login", json = {"email":email, "password":password})
    if not check_for_error(token):token = token.get("token", {'error':''}) 

    return token

def verify_profile_token(token):
    user_id = make_post_request("http://localhost:8080/verify", json = {"token":token})
    if not check_for_error(user_id):user_id = user_id.get("user_id", {'error':''})

    return user_id
