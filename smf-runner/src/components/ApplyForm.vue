<template>
	<form @submit="apply">
		<label>
			Cluster:
			<input v-model="deployment.clusterName" type="text"/>
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
			<button>Apply, bitch.</button>
		</div>
	</form>
</template>
<script>
	import {createComponent, reactive} from "@vue/composition-api";

	export default createComponent({
		setup() {
			const deployment = reactive({
				clusterName: "",
				deploymentFilePath: "",
				plugins: ["istio", "linkrd"]
			});
			return {
				deployment,
				plugins: ["istio", "linkrd"],
				apply() {
					alert("submit");
				}
			};
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

	input[type="text"] {
		margin: 5px 0;
		display: block;
		padding: 5px;
		transition: outline-width 0.3s ease-in-out;
		outline: 0;
		width: 100%;
	}

	input[type="text"]:focus {
		outline: solid 2px #ccc;
	}

	input[type="checkbox"] {
		margin-right: 5px;
		zoom: 2;
		vertical-align: middle;
	}

	.btn-group {
		text-align: center;
	}

	button {
		padding: 10px;
	}
</style>
