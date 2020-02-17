import Vue from 'vue';
import {computed} from '@vue/composition-api';
import axios from 'axios'

const clustersState = Vue.observable({
	clusters: [],
	currentCluster: null
});

function fetchClusters() {
	return axios.get('/api/clusters').then(res => {
		clustersState.clusters = res.data;
	});
}


function fetchCluster(clusterName) {
	return axios.get('/api/clusters/' + clusterName).then(res => {
		clustersState.currentCluster = res.data;
	});
}

export function useClusters() {
	return {
		fetchClusters,
		fetchCluster,
		clusters: computed(() => clustersState.clusters),
		currentCluster: computed(() => clustersState.currentCluster),
	}
}
