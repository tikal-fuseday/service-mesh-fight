const path = require('path');
const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const morgan = require("morgan");

const app = express();
app.use(morgan("combined"));
app.use(cors());

// tell the app to parse HTTP body messages
app.use(bodyParser.json());

require('./api')(app);

app.use(express.static(path.join(__dirname, '../dist')));
app.use('*', express.static(path.join(__dirname, '../dist/index.html')));

app.set("port", process.env.PORT || 3000);
app.set("ip", process.env.IP || "0.0.0.0");

// start the server
app.listen(app.get("port"), app.get("ip"), () => {
	console.log(
		`Service-Mesh-Fight is running.
		\nPlease open in browser: http://${app.get("ip")}:${app.get("port")}`
	);
});
