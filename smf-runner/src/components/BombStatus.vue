<template>
	<div>
		{{bomb}}
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
</style>
