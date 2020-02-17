const {exec} = require("child_process");
const clusters = require('../../clusters');


function getClusters() {
	return Promise.resolve(Object.keys(clusters));
	/*return new Promise((resolve, reject) => {
		exec("kubectl config get-contexts -o=name", function (err, data) {
			if (err) {
				reject(err);
				return;
			}
			resolve(
				data
					.toString()
					.split("\n")
					.filter(Boolean)
			);
		});
	});*/
}

function useKubeCluster(clusterName) {
	return new Promise((resolve, reject) => {
		exec(`kubectl config use-context ${clusterName}`, function (err, data) {
			if (err) {
				reject(err);
				return;
			}
			const result = data.toString();
			if (result.includes('error')) {
				reject(result);
				return;
			}
			resolve(result);
		});
	});
}

async function applyDeployment(clusterName, namespace, deploymentFilePath) {
	await useKubeCluster(clusterName);
	await createNamespace(namespace);

	return new Promise((resolve, reject) => {
		exec(`kubectl apply -f ${deploymentFilePath} -n ${namespace}`, function (err, data) {
			if (err) {
				reject(err);
				return;
			}
			const result = data.toString();
			if (result.includes('error')) {
				reject(result);
				return;
			}
			resolve(result);
		});
	});
}

function getNamespaces() {
	return new Promise((resolve, reject) => {
		exec(`kubectl get ns -o=name`, function (err, data) {
			if (err) {
				reject(err);
				return;
			}
			resolve(
				data
					.toString()
					.split("\n")
					.filter(Boolean)
					.map(name => name.substr('namespace/'.length))
			);
		});
	});
}

async function createNamespace(namespace) {
	const namespaces = await getNamespaces();

	if (namespaces.includes(namespace)) {
		return Promise.resolve();
	}

	return new Promise((resolve, reject) => {
		exec(`kubectl create ns ${namespace}`, function (err, data) {
			if (err) {
				reject(err);
				return;
			}
			const result = data.toString();
			if (result.includes('error')) {
				reject(result);
				return;
			}
			resolve(result);
		});
	});
}

module.exports = {getClusters, applyDeployment};
