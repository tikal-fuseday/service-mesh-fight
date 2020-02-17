<template>
	<div class="bomb-status">
		<p>
			Cluster: <strong>{{bomb.clusterName}}</strong>
		</p>
		<p v-if="bomb.requests">
			Total Requests: <strong>{{bomb.requests}}</strong>
		</p>
		<div class="progress-bar">
			<div class="progress" :style="{width: completed}"></div>
			<strong>{{completed}}</strong>
		</div>
	</div>
</template>
<script>
	import {createComponent} from "@vue/composition-api";
	import {useBomb} from '../compositions/bomb';

	export default createComponent({
		setup() {
			return {
				...useBomb()
			}
		}
	});
</script>
<style scoped>
	.bomb-status {
		width: 100%;
		max-width: 600px;
		margin: 0 auto;
		padding: 10px;
	}
	p {
		font-size: 32px;
	}

	.progress-bar {
		margin-top: 20px;
		position: relative;
		text-align: center;
		width: 100%;
		line-height: 20px;
		z-index: 1;
		background-color: #eee;
		border: 1px solid #ddd;
	}
	.progress {
		position: absolute;
		top: 0;
		left: 0;
		bottom: 0;
		z-index: -1;
		background: rgb(131,58,180);
		background: linear-gradient(90deg, rgba(131,58,180,1) 0%, rgba(252,176,69,1) 100%);
		transition: width 0.2s ease-in-out;
	}
</style>
