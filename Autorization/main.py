from flask import Flask, render_template, request, g, redirect, url_for
import os
import sqlite3 as sq
from FDataBase import FDataBase
from werkzeug.security import generate_password_hash, check_password_hash
from flask_login import LoginManager, login_user, login_required
from UserLogin import UserLogin

# configuration
DATABASE = '/tmp/users.db'
DEBUG = True
SECRET_KEY = 'dkslmr2oiewondo2f;sdw{<NiwNdwqf,'

app = Flask(__name__)
app.config.from_object(__name__)

app.config.update(dict(DATABASE = os.path.join(app.root_path, 'users.db')))

login_manager = LoginManager(app)

@login_manager.user_loader
def load_user(user_id):
    print("load_user")
    return UserLogin().fromDB(user_id, dbase)

def connect_db():
    conn = sq.connect(app.config['DATABASE'])
    conn.row_factory = sq.Row
    return conn


def create_db():
    db = connect_db()
    with app.open_resource('sq_db.sql', mode='r') as f:
        db.cursor().executescript(f.read())
    db.commit()
    db.close()


def get_db():
    if not hasattr(g, 'link_db'):
        g.link_db = connect_db()
    return g.link_db


dbase = None


@app.before_request
def before_request():
    global dbase
    db = get_db()
    dbase = FDataBase(db)


@app.teardown_appcontext
def close_db(error):
    if hasattr(g, 'link_db'):
        g.link_db.close()


@app.route("/")
def Kak_Chto():
    return "/login - логиниться /register - регаться"


@app.route("/login", methods=["POST", "GET"])
def login():
    if request.method == "POST":
        user = dbase.getUserByLogin(request.form['log'])
        if user and check_password_hash(user['psw'], request.form['psw']):
            userlogin = UserLogin().create(user)
            login_user(userlogin)
            return redirect(url_for('login'))
        print("Неверный логин или пароль")
    return render_template('login.html')


@app.route("/register", methods=["POST", "GET"])
def register():
    if request.method == "POST":
        if len(request.form['log']) > 6 and len(request.form['psw']) > 6 and request.form['psw'] == request.form['psw2']:
            hash = generate_password_hash(request.form['psw'])
            res = dbase.addUser(request.form['log'], hash)
            if res:
                print("Пользоваель добавлен")
                return redirect(url_for('register'))
            else:
                print("Ошибка при добавлении в БД")
        else:
            print("Неверно заполнены поля")
    return render_template('register.html')


@app.route("/you")
@login_required
def you():
    return "you."


if __name__ == "__main__":
    app.run(debug=True)