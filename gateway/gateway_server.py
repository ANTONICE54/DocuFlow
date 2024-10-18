from flask import Flask, Blueprint, render_template, request, redirect, session, send_file
from sign_check import *
from auth_functions import *
from user_functions import *
from category_functions import *

docu_flow_app = Flask(__name__)
docu_flow_blueprint = Blueprint('docu_flow', __name__)

docu_flow_app.config.from_object('config.Config')
docu_flow_config = docu_flow_app.config

@docu_flow_blueprint.route('/', methods=['GET'])
def main():
    user_id = None
    verify_response = verify_profile_token(session.get('token', ''))
    if verify_response!='error':
        user_id=verify_response
    return render_template('main.html', have_account=False, user_id=user_id)


@docu_flow_blueprint.route('/sign_in', methods=['GET', 'POST'])
def sign_in():
    verify_response = verify_profile_token(session.get('token', ''))
    if verify_response!='error':
        return redirect('/')

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
    verify_response = verify_profile_token(session.get('token', ''))
    if verify_response!='error':
        return redirect('/')
    
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





@docu_flow_blueprint.route('/account', methods=['GET'])
def account():
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if user_id=='error' or user_info=='error':
       return redirect('/')
    
    return render_template('account.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper())


@docu_flow_blueprint.route('/documents', methods=['GET'])
def documents():
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if user_id=='error' or user_info=='error':
       return redirect('/')
    
    categories_and_sub_categories = get_categories_and_subcategories(user_id)

    if categories_and_sub_categories=='error':
       return redirect('/account')
    
    categories_and_sub_categories = [(category_and_subcategory['name'], [subcategory['name'] for subcategory in category_and_subcategory['subcategory_list']]) for category_and_subcategory in categories_and_sub_categories]
    
    return render_template('documents.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper(), categories_and_sub_categories=categories_and_sub_categories)


@docu_flow_blueprint.route('/create_category', methods=['GET'])
def create_category():
    pass


docu_flow_app.register_blueprint(docu_flow_blueprint)

if __name__ == '__main__':
    docu_flow_app.run(debug=True)
