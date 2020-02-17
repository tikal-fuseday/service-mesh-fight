import Vue from 'vue';
import {computed} from '@vue/composition-api';
import axios from 'axios'

const clustersState = Vue.observable({
	clusters: []
});

function fetchClusters() {
	return axios.get('/api/clusters').then(res => {
		clustersState.clusters = res.data;
	});
}

export function useClusters() {
	return {
		fetchClusters,
		clusters: computed(() => clustersState.clusters)
	}
}
