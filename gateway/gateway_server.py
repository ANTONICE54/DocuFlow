from flask import Flask, Blueprint, render_template, request, redirect, session, send_file

docu_flow_app = Flask(__name__)
docu_flow_blueprint = Blueprint('docu_flow', __name__)

docu_flow_app.config.from_object('config.Config')
docu_flow_config = docu_flow_app.config

@docu_flow_blueprint.route('/', methods=['GET'])
def main():
    return render_template('main.html', have_account=False)


@docu_flow_blueprint.route('/sign_in', methods=['GET'])
def sign_in():
    return render_template('sign_in.html')


@docu_flow_blueprint.route('/sign_up', methods=['GET'])
def sign_up():
    return render_template('sign_up.html')


docu_flow_app.register_blueprint(docu_flow_blueprint)

if __name__ == '__main__':
    docu_flow_app.run(debug=True)