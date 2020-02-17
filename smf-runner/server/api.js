const {getClusters, applyDeployment} = require("./kube");
const axios = require('axios');

module.exports = function (app) {

	app.get("/api/clusters", async function (req, res) {
		const clusters = await getClusters();
		res.status(200).json(clusters);
	});

	app.post("/api/apply", async function (req, res) {
		const body = req.body || {};
		if (!(body.clusterName && body.namespace && body.deploymentFilePath)) {
			res.status(401).json({message: 'missing arguments'});
			return;
		}
		await applyDeployment(body);
		res.status(200).json({status: 'success'});
	});

	app.post("/api/bomb", async function (req, res) {
		axios(`http://${process.env.BOMB_URL}/`, {
			method: 'POST',
			body: {
				time: 1000 * 60 * 60,
				parallels: req.body.parallels || 10
			}
		}).then(({data}) => {
			res.status(200).json(data);
		})
	});

};
