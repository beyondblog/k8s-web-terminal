const URL_ICON_DOCKER = 'assets/svg/docker.svg';
const URL_ICON_MENU    = 'assets/svg/menu.svg';
const URL_ICON_SHARE   = 'assets/svg/share.svg';

// Load the custom app ES6 modules

import DashboardController from 'dashboard/DashboardController'
import DashboardService from 'dashboard/DashboardService'

import {
    ExternalLogger
}
from 'utils/LogDecorator';

let $log = new ExternalLogger();
$log = $log.getInstance("BOOTSTRAP");
$log.debug("Configuring 'Dashboard' module");

// Define the Angular 'Dashboard' module

let moduleName = angular
    .module("dashboard", [])
    .service("dashboardService", DashboardService)
    .controller("DashboardController", DashboardController)
    .config( ($mdIconProvider) => {


      $log.debug( "Configuring $mdIconProvider" );

      // Register `dashboard` iconset & icons for $mdIcon service lookups

      $mdIconProvider
        .defaultIconSet( URL_ICON_DOCKER, 75 )
        .icon('menu' , URL_ICON_MENU, 24)
        .icon('share', URL_ICON_SHARE, 24)
        .icon('docker', URL_ICON_DOCKER, 75);

    })
    .name;

export default moduleName;
