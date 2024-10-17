from flask import Flask, Blueprint, render_template, request, redirect, session, send_file
from sign_check import *
from auth_functions import *

docu_flow_app = Flask(__name__)
docu_flow_blueprint = Blueprint('docu_flow', __name__)

docu_flow_app.config.from_object('config.Config')
docu_flow_config = docu_flow_app.config

@docu_flow_blueprint.route('/', methods=['GET'])
def main():
    return render_template('main.html', have_account=False)


@docu_flow_blueprint.route('/sign_in', methods=['GET', 'POST'])
def sign_in():
    if request.method == 'POST':
        email = request.form.get('email', '')
        password = request.form.get('password', '')

        sign_in_result = check_sign_in({'email': email, 'password': password})

        if sign_in_result != "succes":
            return render_template('sign_in.html', filled_data=[email, password], warning_title=sign_in_result)
  
        login_token = log_in_profile(email, password)

        if login_token == "error":
            return render_template('sign_in.html', filled_data=[email, password], warning_title="Під час входу сталася помилка")
        
        session['token'] = login_token

        return redirect('/')
    return render_template('sign_in.html', filled_data=[])


@docu_flow_blueprint.route('/sign_up', methods=['GET', 'POST'])
def sign_up():
    if request.method == 'POST':
        name = request.form.get('name', '')
        surname = request.form.get('surname', '')
        email = request.form.get('email', '')
        country = request.form.get('country', '')
        password = request.form.get('password', '')
        password_confirmation = request.form.get('password_confirmation', '')

        sign_up_result = check_sign_up({'name': name,'surname': surname, 'email': email, 'country': country, 'password': password, 'password_confirmation': password_confirmation})

        if sign_up_result != "succes":
            return render_template('sign_up.html', filled_data=[name, surname, email, country, password, password_confirmation], warning_title=sign_up_result)
        

        register_token = register_profile(name, surname, email, country, password)

        if register_token == "error":
            return render_template('sign_up.html', filled_data=[name, surname, email, country, password, password_confirmation], warning_title="Щось пішло не так з реєстрацією")
        
        session['token'] = register_token

        return redirect('/')
    return render_template('sign_up.html',  filled_data=[])


docu_flow_app.register_blueprint(docu_flow_blueprint)

if __name__ == '__main__':
    docu_flow_app.run(debug=True)