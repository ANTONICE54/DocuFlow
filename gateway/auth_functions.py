import requests

def register_profile(name, surname, email, country, password):
    for register_attempt in range(3):
        try:
            register_response = requests.post("http://localhost:8080/register", json= {"name":name, "surname":surname, "email":email, "country":country, "password":password})
            register_status = register_response.status_code
            if register_status == 200:
                token = register_response.json()['token']
                return token
        except:
            pass

    return 'error'

def log_in_profile(email, password):
    for login_attempt in range(3):
        try:
            login_response = requests.post("http://localhost:8080/login", json= {"email":email, "password":password})
            login_status = login_response.status_code
            if login_status == 200:
                token = login_response.json()['token']
                return token
        except:
            pass

    return 'error'

def verify_profile_token(token):
    for verify_attempt in range(3):
        try:
            verify_response = requests.post("http://localhost:8080/login", json= {"token":token})
            verify_status = verify_response.status_code
            if verify_status == 200:
                token = verify_response.json()
                print(token)
                return token
        except:
            pass

    return 'error'
