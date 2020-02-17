<template>
	<div>
		<template v-if="cluster">
			<iframe :src="cluster['istio-grafana']"></iframe>
			<iframe :src="cluster['linkerd-grafana']"></iframe>
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
		display: flex;
		flex-direction: column;
	}
</style>
