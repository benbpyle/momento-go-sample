import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { CacheFunction } from "./function";
export class LambdaStack extends cdk.Stack {
    constructor(scope: Construct, id: string) {
        super(scope, id);

        new CacheFunction(this, `Function`);
    }
}
