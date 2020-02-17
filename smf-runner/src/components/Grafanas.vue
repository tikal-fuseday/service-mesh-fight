<template>
	<div>
		<template v-if="cluster">
			<a :href="cluster['istio-grafana']" target="_blank">istio stats</a>
			<a :href="cluster['linkerd-grafana']" target="_blank">linkerd stats</a>
		</template>
	</div>
</template>
<script>
	import {createComponent} from "@vue/composition-api";
	import {useClusters} from '../compositions/clusters';

	export default createComponent({
		props: {
			clusterName: String
		},
		setup(props) {
			const {fetchCluster, currentCluster} = useClusters();
			fetchCluster(props.clusterName);
			return {
				cluster: currentCluster
			}
		}
	});
</script>
<style scoped>
	div {
		text-align: center;
	}

	a {
		color: black;
		font-size: 22px;
		padding: 0 10px;
	}
</style>
