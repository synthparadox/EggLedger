<template>
    <!-- Artifacts -->
    <drop-display
        :item-array="data.artifacts" type="artifact" label-class-list="mission-view-div bg-blue-900"
        label-display-value="Artifacts" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
        :total-drops-count="getTotalCount()"
    ></drop-display>

    <!-- Stones -->
    <drop-display
        :item-array="data.stones" type="stone" label-class-list="mission-view-div bg-fuchsia-900"
        label-display-value="Eggfinity Stones" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
        :total-drops-count="getTotalCount()"
    ></drop-display>

    <!-- Ingredients -->
    <drop-display
        :item-array="data.ingredients" type="ingredient" label-class-list="mission-view-div text-gray-400 bg-darkerer"
        label-display-value="Ingredients" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
        :total-drops-count="getTotalCount()"
    ></drop-display>

    <!-- Stone Fragments -->
    <drop-display
        :item-array="data.stoneFragments" type="stone_fragment" label-class-list="mission-view-div text-gray-400 bg-darkerer"
        label-display-value="Stone Fragments" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
        :total-drops-count="getTotalCount()"
    ></drop-display>

</template>

<script>
    import DropDisplay from './DropDisplay.vue';

    export default {
        components: {
            DropDisplay,
        },
        methods: {
            getTotalCount(){
                const data = this.data;
                return [
                    ...data.artifacts,
                    ...data.stones,
                    ...data.ingredients,
                    ...data.stoneFragments,
                ].reduce((acc, item) => acc + item.count, 0);
            },
        },
        props: {
            data: Object,
            ledgerType: String,
            lifetimeShowPerShip: Boolean,
            lifetimeMissionCount: Number,
            mennoMissionData: Object,
            showExpectedDrops: Boolean,
        },
    }

</script>