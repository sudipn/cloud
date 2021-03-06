<template>
    <div class="follow-panel">
        <span v-if="following" v-on:click="unfollowProject" class="icon">
            <img src="@/assets/icon-heart.svg" class="icon" />
        </span>
        <span v-if="!following" v-on:click="followProject" class="icon">
            <img src="@/assets/icon-heart-gray.svg" class="icon" />
        </span>
        <span v-show="followers > 1">{{ followers }} Follows</span>
        <span v-show="followers <= 1">{{ followers }} Follow</span>
    </div>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import * as ActionTypes from "@/store/actions";

export class ProjectFollowing {
    constructor(public readonly following: boolean, public readonly total: number) {}
}

export default Vue.extend({
    name: "FollowPanel",
    props: {
        project: {
            type: Object,
            required: true,
        },
    },
    data() {
        return {
            action: null,
        };
    },
    computed: {
        following(this: any) {
            return this.action ? this.action.following : this.project.following.following;
        },
        followers(this: any) {
            return this.action ? this.action.total : this.project.following.total;
        },
    },
    methods: {
        async followProject(this: any) {
            await this.$store.dispatch(ActionTypes.PROJECT_FOLLOW, { projectId: this.project.id });
            this.action = new ProjectFollowing(true, this.followersExceptMyself() + 1);
        },
        async unfollowProject(this: any) {
            await this.$store.dispatch(ActionTypes.PROJECT_UNFOLLOW, { projectId: this.project.id });
            this.action = new ProjectFollowing(false, this.followersExceptMyself());
        },
        followersExceptMyself() {
            if (this.project.following.following) {
                return this.project.following.total - 1;
            }
            return this.project.following.total;
        },
    },
});
</script>

<style scoped lang="scss">
@import '../../scss/mixins';

.follow-panel {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 18px;
    font-weight: 500;
    color: #2c3e50;
    padding-top: 24px;
    border-radius: 1px;
    border-top: solid 1px #d8dce0;
    margin-top: 10px;

    @include bp-down($xs) {
        padding-top: 14px;
    }
}
.icon {
    cursor: pointer;
    margin-right: 10px;
    width: 22px;
    height: 20px;
}
</style>
