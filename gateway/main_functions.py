import requests


def make_get_request(url, json={}):
    for attempt in range(3):
        try:
            response = requests.get(url, json=json)
            response_status = response.status_code
            if response_status == 200:
                return response.json()
        except:
            pass

    return 'error'



def make_post_request(url, json={}):
    for attempt in range(3):
        try:
            response = requests.post(url, json=json)
            response_status = response.status_code
            if response_status == 200:
                return response.json()
        except:
            pass

    return 'error'
