from flask import Flask, render_template

app = Flask(__name__)

@app.route('/download')
def logs():
    '''Rendering start page'''

    return render_template('index.html', system_info_text='')


if __name__ == '__main__':
    app.run(debug=True)
