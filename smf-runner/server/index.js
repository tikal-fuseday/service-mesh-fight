const path = require('path');
const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const morgan = require("morgan");
const {getClusters} = require("./kube");

const app = express();
app.use(morgan("combined"));
app.use(cors());

// tell the app to parse HTTP body messages
app.use(bodyParser.json());

app.get("/api/clusters", async function (req, res) {
	const clusters = await getClusters();
	res.status(200).json(clusters);
});

app.post("/api/apply", function (req, res) {
	res.status(200).json({});
});


app.use(express.static(path.join(__dirname, 'dist')));
app.use('*', express.static(path.join(__dirname, 'dist/index.html')));

app.set("port", process.env.PORT || 3000);
app.set("ip", process.env.IP || "0.0.0.0");

// start the server
app.listen(app.get("port"), app.get("ip"), () => {
	console.log(
		`Service-Mesh-Fight is running.
		\nPlease open in browser: http://${app.get("ip")}:${app.get("port")}`
	);
});
