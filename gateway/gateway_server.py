from flask import Flask, Blueprint, render_template, request, redirect, session, send_file
from sign_check import *
from auth_functions import *
from user_functions import *
from category_functions import *
from document_functions import *

docu_flow_app = Flask(__name__)
docu_flow_blueprint = Blueprint('docu_flow', __name__)

docu_flow_app.config.from_object('config.Config')
docu_flow_config = docu_flow_app.config

create_folder(f'document_templates')
create_folder(f'documents')

@docu_flow_blueprint.route('/', methods=['GET'])
def main():
    user_id = None
    verify_response = verify_profile_token(session.get('token', ''))
    print(verify_response)
    if not check_for_error(verify_response):
        user_id=verify_response

    return render_template('main.html', have_account=False, user_id=user_id)


@docu_flow_blueprint.route('/sign_in', methods=['GET', 'POST'])
def sign_in():
    verify_response = verify_profile_token(session.get('token', ''))
    if not check_for_error(verify_response):
        return redirect('/')

    if request.method == 'POST':
        email = request.form.get('email', '')
        password = request.form.get('password', '')

        sign_in_result = check_sign_in({'email': email, 'password': password})

        if sign_in_result != "succes":
            return render_template('sign_in.html', filled_data=[email, password], warning_title=sign_in_result)
  
        login_token = log_in_profile(email, password)

        if check_for_error(login_token):
            return render_template('sign_in.html', filled_data=[email, password], warning_title="Під час входу сталася помилка")
        
        session['token'] = login_token

        return redirect('/')
    return render_template('sign_in.html', filled_data=[])


@docu_flow_blueprint.route('/sign_up', methods=['GET', 'POST'])
def sign_up():
    verify_response = verify_profile_token(session.get('token', ''))
    if not check_for_error(verify_response):
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

        if check_for_error(register_token):
            return render_template('sign_up.html', filled_data=[name, surname, email, country, password, password_confirmation], warning_title="Щось пішло не так з реєстрацією")
        
        session['token'] = register_token

        return redirect('/')
    return render_template('sign_up.html',  filled_data=[])





@docu_flow_blueprint.route('/account', methods=['GET'])
def account():
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if check_for_error(user_id) or check_for_error(user_info):
       return redirect('/')
    
    return render_template('account.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper())


@docu_flow_blueprint.route('/documents', methods=['GET'])
def documents():
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if check_for_error(user_id) or check_for_error(user_info):
       return redirect('/')
    
    categories_and_sub_categories = get_categories_and_subcategories(user_id)

    if categories_and_sub_categories=='error':
       return redirect('/account')
    
    categories_and_sub_categories = sorted([(category_and_subcategory['id'], category_and_subcategory['name'], [subcategory['name'] for subcategory in category_and_subcategory['subcategory_list']]) for category_and_subcategory in categories_and_sub_categories], key=lambda category_elements:category_elements[1])
    
    return render_template('documents.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper(), categories_and_sub_categories=categories_and_sub_categories)


@docu_flow_blueprint.route('/create_category/<string:category_name>', methods=['GET'])
def create_category(category_name):
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if check_for_error(user_id) or check_for_error(user_info):
       return redirect('/')
    
    create_category_response=create_category_function(user_id, category_name)

    if check_for_error(create_category_response):
        categories_and_sub_categories = get_categories_and_subcategories(user_id)

        if check_for_error(categories_and_sub_categories):
            return redirect('/account')
        
        categories_and_sub_categories = sorted([(category_and_subcategory['id'], category_and_subcategory['name'], [subcategory['name'] for subcategory in category_and_subcategory['subcategory_list']]) for category_and_subcategory in categories_and_sub_categories], key=lambda category_elements:category_elements[1])
        
        return render_template('documents.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper(), categories_and_sub_categories=categories_and_sub_categories, warning_title="Сталася помилка при створені категорії")

    return redirect('/documents')


@docu_flow_blueprint.route('/create_subcategory/<int:category_id>/<string:subcategory_name>', methods=['GET'])
def create_subcategory(category_id, subcategory_name):
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if check_for_error(user_id) or check_for_error(user_info):
       return redirect('/')
    
    create_subcategory_response=create_subcategory_function(category_id, subcategory_name)

    if check_for_error(create_subcategory_response):
        categories_and_sub_categories = get_categories_and_subcategories(user_id)

        if check_for_error(categories_and_sub_categories):
            return redirect('/account')
        
        categories_and_sub_categories = sorted([(category_and_subcategory['id'], category_and_subcategory['name'], [subcategory['name'] for subcategory in category_and_subcategory['subcategory_list']]) for category_and_subcategory in categories_and_sub_categories], key=lambda category_elements:category_elements[1])
        
        return render_template('documents.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper(), categories_and_sub_categories=categories_and_sub_categories, warning_title="Сталася помилка при створені підкатегорії")

    return redirect('/documents')
    


@docu_flow_blueprint.route('/create_document', methods=['GET', 'POST'])
def create_document():
    user_id = verify_profile_token(session.get('token', ''))
    user_info = get_user_info(user_id)

    if check_for_error(user_id) or check_for_error(user_info):
       return redirect('/')
    
    if request.method == 'POST':
        document_template = request.files['template_document']
        create_request = request.form['create_request']
        document_template_text = get_document_text((document_template.filename, document_template.stream, document_template.mimetype))
        if not check_for_error(document_template_text):
            document_name = ''.join(document_template.filename.split('.')[:-1])
            template_document_name = save_document_template(user_id, document_name, document_template_text)
            document_text = create_document_by_template(user_id, template_document_name, create_request)

            save_document(user_id, document_name, document_text)
    
    return render_template('create_document.html', avatar_name=f"{user_info['name'][0]}{user_info['surname'][0]}".upper())



docu_flow_app.register_blueprint(docu_flow_blueprint)

if __name__ == '__main__':
    docu_flow_app.run(debug=True)
