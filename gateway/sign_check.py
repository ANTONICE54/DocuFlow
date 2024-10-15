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
        return f"Пітвердження паролю не співпадає"
        
    return "succes"
    