import requests

def register_profile(name, surname, email, country, password):
    register_status = 0

    for register_attempt in range(3):
        try:
            response = requests.get(url, params= params, timeout=20)
        except:
            pass
        register_status = response.status_code

def log_in_profile(email, password):
    pass

def reload_profile_token(email, password):
    pass