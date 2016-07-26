app.controller('XeditableCtrl', ['$scope', '$filter', '$http', 'editableOptions', 'editableThemes', 
  function($scope, $filter, $http, editableOptions, editableThemes){
    editableThemes.bs3.inputClass = 'input-sm';
    editableThemes.bs3.buttonsClass = 'btn-sm';
    editableOptions.theme = 'bs3';

    $scope.html5 = {
      email: 'email@example.com',
      tel: '123-45-67',
      number: 29,
      range: 10,
      url: 'http://example.com',
      search: 'blabla',
      color: '#6a4415',
      date: null,
      time: '12:30',
      datetime: null,
      month: null,
      week: null
    };

    $scope.user = {
    	name: 'awesome',
    	desc: 'Awesome user \ndescription!',
      status: 2,
      agenda: 1,
      remember: false
    };

    $scope.statuses = [
      {value: 1, text: 'status1'},
      {value: 2, text: 'status2'},
      {value: 3, text: 'status3'}
    ];

    $scope.agenda = [
      {value: 1, text: 'male'},
      {value: 2, text: 'female'}
    ];

    $scope.showStatus = function() {
      var selected = $filter('filter')($scope.statuses, {value: $scope.user.status});
      return ($scope.user.status && selected.length) ? selected[0].text : 'Not set';
    };

    $scope.showAgenda = function() {
      var selected = $filter('filter')($scope.agenda, {value: $scope.user.agenda});
      return ($scope.user.agenda && selected.length) ? selected[0].text : 'Not set';
    };

    // editable table
    $scope.users = [
      {id: 1, name: 'awesome user1', status: 2, group: 4, groupName: 'admin'},
      {id: 2, name: 'awesome user2', status: undefined, group: 3, groupName: 'vip'},
      {id: 3, name: 'awesome user3', status: 2, group: null}
    ];

    $scope.things = []
    $scope.views = []

    $scope.groups = [];
    $scope.loadGroups = function() {
      return $scope.groups.length ? null : $http.get('api/groups').success(function(data) {
        $scope.groups = data;
      });
    };

    $scope.showGroup = function(user) {
      if(user.group && $scope.groups.length) {
        var selected = $filter('filter')($scope.groups, {id: user.group});
        return selected.length ? selected[0].text : 'Not set';
      } else {
        return user.groupName || 'Not set';
      }
    };

    $scope.showStatus = function(user) {
      var selected = [];
      if(user && user.status) {
        selected = $filter('filter')($scope.statuses, {value: user.status});
      }
      return selected.length ? selected[0].text : 'Not set';
    };

    $scope.checkName = function(data, id) {
      if (id === 2 && data !== 'awesome') {
        return "Username 2 should be `awesome`";
      }
    };

    $scope.saveUser = function(data, id) {
      //$scope.user not updated yet
      angular.extend(data, {id: id});
      console.dir(data)
      // return $http.post('api/saveUser', data);
    };

    // remove user444444
    $scope.removeThing = function(index) {
      $scope.things.splice(index, 1);
    };

    $scope.addView = function() {
      inserted = {
        id: "",
        name: '',
        room: '',
        floor: 0,
        zone_name: '',
        thing:[]
      };
      $scope.views.push(inserted);
    };

    // add thing
    $scope.addThing = function() {
      $scope.inserted = {
        id: "",
        name: '',
        type: '',
        display_topic: '',
        control_topic: '',
        ui_element:'sensor',
        max_value:0,
        min_value:0,
        value:null,
        unit:"",
        updated_at:"0001-01-01T00:00:00Z",
        prop_field_ui:""
      };
      $scope.things.push($scope.inserted);
    };



    //$scope.loadProject = function(){
    //    $http.get('/greenhome/api/project/57573834554efc2c77b59f97').
    //    then(function (response) {
    //        $scope.project = response.data
    //    }, function (response) {
    //        // called asynchronously if an error occurs
    //        // or server returns response with an error status.
    //    });
    //}
    $scope.loadProjects = function(){
        $http.get('/greenhome/api/projects').
        then(function (response) {
            $scope.projects = response.data
        }, function (response) {
            // called asynchronously if an error occurs
            // or server returns response with an error status.
        });
    }
    $scope.saveProject = function(){
      $http.post("/greenhome/api/project",JSON.stringify($scope.project)).
      then(function (response) {
            alert("Project saved")
        }, function (response) {

            // called asynchronously if an error occurs
            // or server returns response with an error status.
        });

    }
    $scope.loadProjects()
    $scope.loadView = function(view) {
      //console.dir(view)
      $scope.things = view.thing
    };
    $scope.loadProject = function(project) {
      //console.dir(project)
      $scope.project = project
      $scope.views = project.view
      $scope.things = []
    };
    $scope.showStruct = function(){
      console.dir($scope.project)
    }


}]);
