const axios = require('axios');
const {getClusters, applyDeployment} = require("./kube");
const clusters = require('../../clusters');

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
		const cluster = clusters[req.body.clusterName];
		if(!cluster) {
			res.status(401).json({message: 'cluster name does not exist'});
			return;
		}

		axios(`${cluster.bombServiceUrl}/api/bomb`, {
			method: 'POST',
			body: {
				time: 1000 * 60 * 60,
				parallels: req.body.parallels || 10
			}
		}).then(({data}) => {
			// should be {bombId: number}
			res.status(200).json(data);
		});
	});

	app.post("/api/bomb/:bombId/status", async function (req, res) {
		const cluster = clusters[req.body.clusterName];
		if(!cluster) {
			res.status(401).json({message: 'cluster name does not exist'});
			return;
		}

		axios(`${cluster.bombServiceUrl}/api/bomb/${req.params.bombId}/status`, {
			method: 'POST',
			body: {
				time: 1000 * 60 * 60,
				parallels: req.body.parallels || 10
			}
		}).then(({data}) => {
			// should be {bombId: number}
			res.status(200).json(data);
		});
	});

};
