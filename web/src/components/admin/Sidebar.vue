<script setup>
import { onMounted, ref } from 'vue';
import { useUserStore } from '@/stores/user.store';

const user = useUserStore();
const userIsAdmin = ref(false);

onMounted(() => {
	userIsAdmin.value = user.current?.is_admin;
	feather.replace()
});

</script>

<template>
	<nav id="sidebar" class="sidebar js-sidebar">
		<div class="sidebar-content js-simplebar" data-simplebar="init">
			<div class="sidebar-brand">
				<span class="align-middle">AI Avatar</span>
			</div>

			<ul class="sidebar-nav" v-if="user.avatar">

				<li class="sidebar-item">
					<router-link :to="{ name: 'admin', params: { avatar_id: user.avatar?.ID } }" class="sidebar-link">
						<i class="align-middle" data-feather="message-circle"></i> <span class="align-middle">Chat</span>
					</router-link>
				</li>

				<li class="sidebar-header">Configure</li>

				<li class="sidebar-item">
					<router-link :to="{ name: 'personality', params: { avatar_id: user.avatar?.ID } }" class="sidebar-link">
						<i class="align-middle" data-feather="meh"></i> <span class="align-middle">Personality</span>
					</router-link>
				</li>

				<li class="sidebar-item">
					<router-link :to="{ name: 'models', params: { avatar_id: user.avatar?.ID } }" class="sidebar-link">
						<i class="align-middle" data-feather="codesandbox"></i> <span class="align-middle">Models</span>
					</router-link>
				</li>

				<li class="sidebar-item">
					<router-link :to="{ name: 'tools', params: { avatar_id: user.avatar?.ID } }" class="sidebar-link">
						<i class="align-middle" data-feather="tool"></i> <span class="align-middle">Tools</span>
					</router-link>
				</li>

				<li class="sidebar-item">
					<router-link :to="{ name: 'agents', params: { avatar_id: user.avatar?.ID } }" class="sidebar-link">
						<i class="align-middle" data-feather="layers"></i> <span class="align-middle">Agents</span>
					</router-link>
				</li>

				<li class="sidebar-item">
					<router-link :to="{ name: 'plugins', params: { avatar_id: user.avatar?.ID } }" class="sidebar-link">
						<i class="align-middle" data-feather="package"></i> <span class="align-middle">Plugins</span>
					</router-link>
				</li>

			</ul>

			<ul class="sidebar-nav" v-else>
				<li class="sidebar-item">
					<router-link :to="{ name: 'create-avatar' }" class="sidebar-link">
						<i class="align-middle" data-feather="message-circle"></i> <span class="align-middle">Create Avatar</span>
					</router-link>
				</li>
			</ul>

			<ul class="sidebar-nav" v-if="userIsAdmin">
				<li class="sidebar-header">Administration</li>

				<li class="sidebar-item">
					<router-link :to="{ name: 'users' }" class="sidebar-link">
						<i class="align-middle" data-feather="users"></i> <span class="align-middle">Manage Users</span>
					</router-link>
				</li>
			</ul>

		</div>
	</nav>
</template>

<style scoped>
.sidebar .card {
	box-shadow: inset 0 3px 3px rgba(0, 0, 0, .1);
	border: 2px solid #ddd;
	border-radius: 6px;
	padding: 15px;
}

.list-group-item {
	background-color: transparent;
	overflow: hidden;
}
</style>
