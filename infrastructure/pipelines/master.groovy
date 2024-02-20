microservice_pipeline([
    'service': [
        'main': 'auth-faker',
        'name': 'auth-faker',
        'environments': [
            'deploy_development_environment': false,
            'deploy_stage_environment': false,
            'deploy_production_environment': false,
            'cross_env': [
                'stack_name_format': 'environment_as_prefix',
                'deploy_with_terraform': true
            ]
        ]
    ],
    'pipeline': [
        'architectures': ['x86_64', 'arm64'],
        'use_semantic_release': true,
    ],
    'code_analysis': [
        'enabled': false
    ]
])
