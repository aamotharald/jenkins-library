import groovy.transform.Field

@Field String STEP_NAME = getClass().getName()
@Field String METADATA_FILE = 'metadata/abapLandscapePortalUpdateAddOnProduct.yaml'

void call(Map parameters = [:]) {
    List credentials = [
        [type: 'token', id: 'landscapePortalAPICredentialsId', env: ['PIPER_landscapePortalAPIServiceKey']]
    ]
    piperExecuteBin(parameters, STEP_NAME, METADATA_FILE, credentials)
}