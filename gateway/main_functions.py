import requests
import traceback
import os 

def make_get_request(url, json={}, params={}, files={}, headers={}):
    error = ""
    for attempt in range(3):
        try:
            response = requests.get(url, json=json, params=params, files=files, headers=headers)
            response_status = response.status_code
            if response_status == 200:
                return response.json()
            else:
                error=str(response.text)

        except:
            error=traceback.format_exc()

    return {"error":error}



def make_post_request(url, json={}, params={}, files={}, headers={}):
    error=''
    for attempt in range(3):
        try:
            response = requests.post(url, json=json, params=params, files=files, headers=headers)
            response_status = response.status_code
            if response_status == 200:
                return response.json()
            else:
                error=str(response.text)
        except:
            error=traceback.format_exc()

    return {"error":error}


def check_for_error(value):
    if type(value)==dict and 'error' in value:
       return True
    return False


def create_folder(folder_name):
    if not os.path.exists(folder_name):
        os.makedirs(folder_name)


def get_unique_name(folder, name, extension):
    counter, unique_name = 1, name
    
    while os.path.exists(f'{folder}/{unique_name}{extension}'):
        unique_name = f"{name}_{counter}"
        counter += 1
    
    return unique_name