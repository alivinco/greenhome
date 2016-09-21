app.controller('ProjectManagerCtrl', ['$scope', '$filter', '$http', 'editableOptions', 'editableThemes',
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

    $scope.things = [];
    $scope.views = [];
    $scope.domains = [];
    $scope.projects = [];
    $scope.groups = [];
    $scope.project = null;
    $scope.view = null;
    $scope.domain = null;
    $scope.ui_elements = ["binary_switch","moment_button","sensor","short_text","level_slider"]
    $scope.username = userName;

    $scope.addProject = function() {
      project = {
        id: "",
        name: '',
        type: 'mob_html',
        domain: $scope.domain.Id,
        comments: '',
        view:[]
      };
      $scope.projects.push(project);
      $scope.project = project
    };

    $scope.deleteProject = function(projectId){
      $http.delete('/greenhome/api/project/'+projectId).
        then(function (response) {
            $scope.loadProjects()
            alert("Project was deleted.")
        }, function (response) {
            // called asynchronously if an error occurs
            // or server returns response with an error status.
            alert(response)
        });
    }

    $scope.addView = function() {
      if ($scope.project){
        inserted = {
        id: "",
        name: '',
        room: '',
        floor: 0,
        zone_name: '',
        thing:[]
        };
        $scope.views.push(inserted);
        $scope.view = inserted
      }else{
        alert("No active project has been selected . Select project first")
      }

    };

    // remove view
    $scope.removeView = function(index) {
      $scope.views.splice(index, 1);
    };

    // add thing
    $scope.addThing = function() {
      if ($scope.view){
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
      }else{
        alert("No active view has been selected . Select view first")
      }

    };

    // remove thing
    $scope.removeThing = function(index) {
      $scope.things.splice(index, 1);
    };


    $scope.loadProjects = function(){
        $http.get('/greenhome/api/projects').
        then(function (response) {
            $scope.projects = response.data
        }, function (response) {
            // called asynchronously if an error occurs
            // or server returns response with an error status.
        });
    }

    $scope.loadDomains = function(){
            $http.get('/greenhome/api/domains').
            then(function (response) {
                $scope.domains = response.data
            }, function (response) {
                // called asynchronously if an error occurs
                // or server returns response with an error status.
            });
        }

    $scope.setDomain = function(domain){
          $scope.domain = domain
          if ($scope.domains){
            $http.post("/greenhome/api/session",{SessionDomain:domain.Id}).
            then(function (response) {
                $scope.loadProjects()
//                alert("Project saved")

              }, function (response) {

                  // called asynchronously if an error occurs
                  // or server returns response with an error status.
              });
          }else{
            alert("Nothing to save , please select project first.")
          }
        }

    $scope.saveProject = function(){
      if ($scope.project){
        $http.post("/greenhome/api/project",JSON.stringify($scope.project)).
        then(function (response) {
            $scope.loadProjects()
            alert("Project saved")

          }, function (response) {

              // called asynchronously if an error occurs
              // or server returns response with an error status.
          });
      }else{
        alert("Nothing to save , please select project first.")
      }
    }
    $scope.loadView = function(view) {
      //console.dir(view)
      $scope.things = view.thing
      $scope.view = view
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
    $scope.loadProjects()
    $scope.loadDomains()

}]);
