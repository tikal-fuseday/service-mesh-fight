<template>
	<div id="app">
		<ApplyForm/>
	</div>
</template>

<script>
	import ApplyForm from "./components/ApplyForm.vue";

	export default {
		name: "App",
		components: {
			ApplyForm
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
	}
</style>
