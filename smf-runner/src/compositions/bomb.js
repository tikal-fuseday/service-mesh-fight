import {computed, reactive} from '@vue/composition-api';
import axios from 'axios'
import Vue from 'vue';

const STORAGE_KEY = 'smf_bomb';

export const bomb = Vue.observable(localStorage.getItem(STORAGE_KEY) ? JSON.parse(localStorage.getItem(STORAGE_KEY)) : {
	clusterName: '',
	parallels: 100,
	bombId: '',
	status: '',
	completed: 0,
	requests: 0,
});

export function createBomb(clusterName, testUrl, parallels = 100) {
	bomb.clusterName = clusterName;
	bomb.parallels = parallels;
	localStorage.setItem(STORAGE_KEY, JSON.stringify(bomb));
	return axios.post(`/api/bomb/${clusterName}`, {url: testUrl}).then(res => res.data)
		.then(data => {
			if (data.bombId) {
				bomb.bombId = data.bombId;
				localStorage.setItem(STORAGE_KEY, JSON.stringify(bomb));
				checkBomb();
			}
		});
}

export function checkBomb() {
	if (!bomb.bombId || !bomb.clusterName) {
		return;
	}
	return axios.get(`/api/bomb/${bomb.clusterName}/${bomb.bombId}/status`)
		.then(res => res.data)
		.then(({status, completed, requests}) => {
			bomb.status = status;
			bomb.completed = completed;
			bomb.requests = requests;
			localStorage.setItem(STORAGE_KEY, JSON.stringify(bomb));
			setTimeout(checkBomb, 2000);
		});
}

export function useBomb() {
	return {
		bomb: reactive(bomb),
		completed: computed(() => (bomb.completed).toFixed(2) + '%')
	}
}
