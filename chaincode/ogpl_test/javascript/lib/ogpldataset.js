/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Contract } = require('fabric-contract-api');

class OGPLDataset extends Contract {

    async initLedger(ctx) {
        console.info('============= START : Initialize Ledger ===========');
        const datasets = [
            {
                name: 'datasetA',
                size: '1400000',
                extension: 'csv',
                owner: 'kha',
                createdtimestamp: '123',
                lastmodifiedtimestamp: '123',
                checksum: 'abc'
            },
            {
                name: 'datasetB',
                size: 4400000,
                extension: 'xml',
                owner: 'anonymous',
                createdtimestamp: 123,
                lastmodifiedtimestamp: 123,
                checksum: 'abc'
            },
            {
                name: 'datasetC',
                size: 2900000,
                extension: 'xls',
                owner: 'anonymous',
                createdtimestamp: 123,
                lastmodifiedtimestamp: 123,
                checksum: 'abc'
            },
            {
                name: 'datasetD',
                size: 5200000,
                extension: 'txt',
                owner: 'anonymous',
                createdtimestamp: 123,
                lastmodifiedtimestamp: 123,
                checksum: 'abc'
            },
        ];

        for (let i = 0; i < datasets.length; i++) {
            datasets[i].docType = 'dataset';
            await ctx.stub.putState('DATASET' + i, Buffer.from(JSON.stringify(datasets[i])));
            console.info('Added <--> ', datasets[i]);
        }
        console.info('============= END : Initialize Ledger ===========');
    }

    async queryDataset(ctx, datasetNumber) {
        const datasetAsBytes = await ctx.stub.getState(datasetNumber); // get the dataset from chaincode state
        if (!datasetAsBytes || datasetAsBytes.length === 0) {
            throw new Error(`${datasetNumber} does not exist`);
        }
        console.log(datasetAsBytes.toString());
        return datasetAsBytes.toString();
    }

    async createDataset(ctx, datasetNumber,name, size, extension , owner, createdtimestamp, checksum) {
        console.info('============= START : Create Dataset ===========');
        var lastmodifiedtimestamp = createdtimestamp;
        const dataset = 
            {
                name,
                size,
                extension,
                owner,
                createdtimestamp,
                lastmodifiedtimestamp,
                checksum
            }

        await ctx.stub.putState(datasetNumber, Buffer.from(JSON.stringify(dataset)));
        console.info('============= END : Create Dataset ===========');
    }

    async queryAllDatasets(ctx) {
        const startKey = 'DATASET0';
        const endKey = 'DATASET999';

        const iterator = await ctx.stub.getStateByRange(startKey, endKey);

        const allResults = [];
        while (true) {
            const res = await iterator.next();

            if (res.value && res.value.value.toString()) {
                console.log(res.value.value.toString('utf8'));

                const Key = res.value.key;
                let Record;
                try {
                    Record = JSON.parse(res.value.value.toString('utf8'));
                } catch (err) {
                    console.log(err);
                    Record = res.value.value.toString('utf8');
                }
                allResults.push({ Key, Record });
            }
            if (res.done) {
                console.log('end of data');
                await iterator.close();
                console.info(allResults);
                return JSON.stringify(allResults);
            }
        }
    }

    async changeDatasetOwner(ctx, datasetNumber, newOwner) {
        console.info('============= START : changeDatasetOwner ===========');

        const datasetAsBytes = await ctx.stub.getState(datasetNumber); // get the dataset from chaincode state
        if (!datasetAsBytes || datasetAsBytes.length === 0) {
            throw new Error(`${datasetNumber} does not exist`);
        }
        const dataset = JSON.parse(datasetAsBytes.toString());
        dataset.owner = newOwner;

        await ctx.stub.putState(datasetNumber, Buffer.from(JSON.stringify(dataset)));
        console.info('============= END : changeDatasetOwner ===========');
    }

    async updateDatasetChecksum(ctx, datasetNumber, newChecksum) {
        console.info('============= START : changeDatasetOwner ===========');

        const datasetAsBytes = await ctx.stub.getState(datasetNumber); // get the dataset from chaincode state
        if (!datasetAsBytes || datasetAsBytes.length === 0) {
            throw new Error(`${datasetNumber} does not exist`);
        }
        const dataset = JSON.parse(datasetAsBytes.toString());
        dataset.checksum = newChecksum;

        await ctx.stub.putState(datasetNumber, Buffer.from(JSON.stringify(dataset)));
        console.info('============= END : updateDatasetChecksum ===========');
    }

}

module.exports = OGPLDataset;
