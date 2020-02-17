import Vue from 'vue';
import {reactive} from '@vue/composition-api';
import axios from 'axios'
import {createBomb} from './bomb';

const STORAGE_KEY = 'smf_deployment';

const plugins = ["istio", "linkerd"];
const deployment = Vue.observable(localStorage.getItem(STORAGE_KEY) ? JSON.parse(localStorage.getItem(STORAGE_KEY)) : {
	clusterName: "",
	namespace: 'istio',
	testUrl: '',
	deploymentFilePath: "",
	plugins: [...plugins],
	loading: false
});

function apply() {
	localStorage.setItem(STORAGE_KEY, JSON.stringify(deployment));
	deployment.loading = true;
	return axios.post('/api/apply', deployment)
		.then(() => {
			return createBomb(deployment.clusterName, deployment.testUrl)
		})
		.finally(() => {
			deployment.loading = false;
		})
}

export function useDeployment() {
	return {
		plugins,
		deployment: reactive(deployment),
		apply
	}
}
