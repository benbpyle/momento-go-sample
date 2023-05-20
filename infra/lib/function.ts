import * as sqs from "aws-cdk-lib/aws-sqs";
import { IQueue, QueueEncryption } from "aws-cdk-lib/aws-sqs";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import { TableEncryption } from "aws-cdk-lib/aws-dynamodb";
import * as kms from "aws-cdk-lib/aws-kms";
import * as path from "path";
import {
    CfnOutput,
    Duration,
    Fn,
    NestedStack,
    RemovalPolicy,
    Tags,
} from "aws-cdk-lib";
import { SqsEventSource } from "aws-cdk-lib/aws-lambda-event-sources";
import * as golambda from "@aws-cdk/aws-lambda-go-alpha";
import { IFunction } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";
import { Secret } from "aws-cdk-lib/aws-secretsmanager";

export class CacheFunction extends Construct {
    constructor(scope: Construct, id: string) {
        super(scope, id);
        const version = Math.round(new Date().getTime() / 1000).toString();

        let table = new dynamodb.Table(this, id, {
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
            removalPolicy: RemovalPolicy.DESTROY,
            partitionKey: { name: "pk", type: dynamodb.AttributeType.STRING },
            pointInTimeRecovery: false,
            tableName: "SampleLookupTable",
        });

        let func = new golambda.GoFunction(this, `CacheFunction`, {
            entry: path.join(__dirname, "../../resources/router"),
            functionName: "cache-function",
            timeout: Duration.seconds(30),
            environment: {
                DD_FLUSH_TO_LOG: "true",
                DD_TRACE_ENABLED: "true",
                LOG_LEVEL: "debug",
                CACHE_NAME: "sample-cache",
                TABLE_NAME: "SampleLookupTable",
                IS_LOCAL: "false",
            },
        });

        Tags.of(func).add("version", version);
        table.grantReadWriteData(func);
        const s = Secret.fromSecretNameV2(this, "Secrets", "mo-cache-token");
        s.grantRead(func);
    }
}
