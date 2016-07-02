/**
 * Users DataService
 * Uses embedded, hard-coded data model; acts asynchronously to simulate
 * remote data service call(s).
 *
 * @returns {{loadAll: Function}}
 * @constructor
 */
function DashboardService($log, $http) {
    var users = [{
        Names: ['Lia Lugo'],
        Id: 'beyond',
    }];

    $log = $log.getInstance("DashboardService");
    $log.debug("instanceOf() ");

    // Promise-based API
    return {
        loadAll: function() {
            $log.debug("loadAll()");
            return $http.get('/api/nodes');
        }
    };
}

export
default ['$log', '$http', DashboardService];
