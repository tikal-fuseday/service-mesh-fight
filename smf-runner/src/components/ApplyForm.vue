<template>
	<form @submit.prevent="apply">
		<label>
			Cluster:
			<select v-model="deployment.clusterName">
				<option v-for="cluster of clusters" :key="cluster" :value="cluster">{{cluster}}</option>
			</select>
		</label>
		<label>
			Namespace:
			<input v-model="deployment.namespace" type="text"/>
		</label>
		<label>
			Test Url:
			<input v-model="deployment.testUrl" type="text"/>
		</label>
		<label>
			Deployment File Path:
			<input v-model="deployment.deploymentFilePath" type="text"/>
		</label>
		<div class="plugins">
			Plugins:
			<label v-for="plugin of plugins" :key="plugin">
				<input type="checkbox" v-model="deployment.plugins" :value="plugin"/>
				<span>{{ plugin }}</span>
			</label>
		</div>
		<div class="btn-group">
			<LoadSpinner v-if="deployment.loading" />
			<button v-else>Apply, bitch.</button>
		</div>
	</form>
</template>
<script>
	import {createComponent} from "@vue/composition-api";
	import {useDeployment} from '../compositions/deployment';
	import {useClusters} from '../compositions/clusters';
	import LoadSpinner from './LoadSpinner';

	export default createComponent({
		components: {
			LoadSpinner,
		},
		setup() {
			const {clusters, fetchClusters} = useClusters();
			fetchClusters();
			return {
				clusters,
				...useDeployment()
			}
		}
	});
</script>
<style scoped>
	form {
		width: 100%;
		max-width: 400px;
		margin: 0 auto;
		padding: 10px;
	}

	label {
		display: block;
		padding: 10px;
	}

	label span {
		vertical-align: middle;
	}

	.plugins {
		padding: 10px;
	}

	input[type="text"], select {
		margin: 5px 0;
		display: block;
		padding: 5px;
		transition: outline-width 0.3s ease-in-out;
		outline: 0;
		width: 100%;
		font-size: 22px;
	}

	input[type="text"]:focus, select:focus {
		outline: solid 2px #ccc;
	}

	input[type="checkbox"] {
		margin-right: 5px;
		zoom: 2;
		vertical-align: middle;
	}

	.btn-group {
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	button {
		padding: 10px;
	}
</style>
