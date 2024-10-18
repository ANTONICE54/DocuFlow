from main_functions import *

def register_profile(name, surname, email, country, password):
    token = make_post_request("http://localhost:8080/register", json = {"name":name, "surname":surname, "email":email, "country":country, "password":password})
    if type(token)!=str: token = token.get("token", 'error') 

    return token

def log_in_profile(email, password):
    token = make_post_request("http://localhost:8080/login", json = {"email":email, "password":password})
    if type(token)!=str: token = token.get("token", 'error') 

    return token

def verify_profile_token(token):
    user_id = make_post_request("http://localhost:8080/verify", json = {"token":token})
    if type(user_id)!=str: user_id = user_id.get("user_id", 'error') 

<<<<<<< HEAD
    return 'error'
=======
    return user_id
>>>>>>> baab450 (	modified:   gateway/auth_functions.py)
