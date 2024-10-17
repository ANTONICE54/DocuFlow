import requests

def check_sign_in(sign_data):
    for data_name, data_value in sign_data.items():
        if data_value=='':
            return f"{data_name} є порожній"
        
    return "succes"
    

def check_sign_up(sign_data):
    for data_name, data_value in sign_data.items():
        if data_value=='':
            return f"{data_name} є порожній"
        
    if sign_data["password"] != sign_data["password_confirmation"]:
        return f"{data_name} пітвердження паролю не співпадає"
    data = {
        "name": sign_data["name"],
        "surname": sign_data["surname"],
        "email": sign_data["email"],
        "country": sign_data["country"],
        "password":sign_data["password"]
    }

    return "succes"
    