<template>
    <!-- Artifacts -->
    <drop-display
        :if-count="data.artifactCount" :item-array="data.artifacts" type="artifact" label-class-list="mission-view-div bg-blue-900"
        label-display-value="Artifacts" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :use-gifs-for-rarity="useGifsForRarity" :af-rarity-class="afRarityClass" :af-rarity-text="afRarityText"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
    ></drop-display>

    <!-- Stones -->
    <drop-display
        :if-count="data.stoneCount" :item-array="data.stones" type="stone" label-class-list="mission-view-div bg-fuchsia-900"
        label-display-value="Eggfinity Stones" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
    ></drop-display>

    <!-- Ingredients -->
    <drop-display
        :if-count="data.ingredientCount" :item-array="data.ingredients" type="ingredient" label-class-list="mission-view-div text-gray-400 bg-darkerer"
        label-display-value="Ingredients" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
    ></drop-display>

    <!-- Stone Fragments -->
    <drop-display
        :if-count="data.stoneFragmentCount" :item-array="data.stoneFragments" type="stone_fragment" label-class-list="mission-view-div text-gray-400 bg-darkerer"
        label-display-value="Stone Fragments" :ledger-type="ledgerType"
        :lifetime-show-per-ship="lifetimeShowPerShip" :lifetime-mission-count="lifetimeMissionCount"
        :menno-mission-data="mennoMissionData" :show-expected-drops="showExpectedDrops"
    ></drop-display>

</template>

<script>
    import DropDisplay from './DropDisplay.vue';

    export default {
        components: {
            DropDisplay,
        },
        props: {
            useGifsForRarity: Boolean,
            data: Object,
            ledgerType: String,
            lifetimeShowPerShip: Boolean,
            lifetimeMissionCount: Number,
            afRarityClass: Function,
            afRarityText: Function,
            mennoMissionData: Object,
            showExpectedDrops: Boolean,
        },
    }

</script>