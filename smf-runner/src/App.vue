<template>
	<div id="app">
		<BombStatus v-if="bomb.bombId"/>
		<ApplyForm v-else/>
	</div>
</template>

<script>
	import BombStatus from "./components/BombStatus.vue";
	import ApplyForm from "./components/ApplyForm.vue";
	import {checkBomb, useBomb} from './compositions/bomb';

	export default {
		name: "App",
		components: {
			BombStatus,
			ApplyForm
		},
		setup() {
			checkBomb();

			return {
				...useBomb()
			}
		}
	};
</script>

<style>
	html,
	body,
	* {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
	}

	html,
	body {
		height: 100%;
		font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
		font-size: 13px;
		background: rgb(69,184,187);
		background: linear-gradient(180deg, rgba(69,184,187,1) 0%, rgba(255,255,255,1) 100%);
	}

	#app {
		height: 100%;
		background: url("./assets/service-mesh.png") no-repeat bottom;
	}
</style>
