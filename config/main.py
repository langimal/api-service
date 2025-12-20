from flask import Flask, jsonify
from api_service.config import Config
from api_service.models import db

app = Flask(__name__)
app.config.from_object(Config)
db.init_app(app)

from api_service.views import api

app.register_blueprint(api, url_prefix='/api')

@app.errorhandler(404)
def not_found(e):
    return jsonify({"error": "Not found"}), 404

@app.errorhandler(500)
def internal_server_error(e):
    return jsonify({"error": "Internal server error"}), 500

if __name__ == '__main__':
    app.run(debug=True)