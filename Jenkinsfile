@Library('dst-shared@master') _

dockerBuildPipeline {
    repository = "cray"
    imagePrefix = "cray"
    app = "spire-bundle"
    name = "spire-bundle"
    description = "Service for serving spire server information and certificate bundle"
        product = "csm"
    githubPushRepo = "Cray-HPE/spire-servebundle"
    githubPushBranches = /(release\/.*|master)/
}
