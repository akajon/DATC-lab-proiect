var mssql = require('mssql');

console.log('Running SQL Query');

var config = {
    user: 'CloudSA35efb96b',
    password: '22.dejlol',
    server: 'proiectdatc.database.windows.net', 
    database: 'city_danger_alert',
    port: 1433
};

var dangers = [0, 0, 0 ,0, 0]

mssql.connect(config, function(err) {
    // ... error checks
    
    if(err) {
        console.log('Connection broke');
        return console.log(err);
    }

    var request = new mssql.Request();
    request.stream = true; // You can set streaming differently for each request
    request.query("select * from dbo.dangers"); // or request.execute(procedure);

    request.on('recordset', function(columns) {
        // Emitted once for each recordset in a query
    });

    request.on('row', function(row) {
        // Emitted for each row in a recordsets
        dangers[row.grade - 1]++;
        console.log(row);
    });

    request.on('error', function(err) {
        // May be emitted multiple times
        console.log('Query broke');
        console.log(err);
    });

    request.on('done', function(returnValue) {
        // Always emitted as the last one
        console.log('Completed Query');
        console.log('There are ' + dangers[0] + ' dangers with grade 1');
        console.log('There are ' + dangers[1] + ' dangers with grade 2');
        console.log('There are ' + dangers[2] + ' dangers with grade 3');
        console.log('There are ' + dangers[3] + ' dangers with grade 4');
        console.log('There are ' + dangers[4] + ' dangers with grade 5');
        process.exit(0);
    });
});

mssql.on('error', function(err) {
    // ... error handler
    console.log('Could not connect to SQL DB');
    process.exit(1);
});

process.on('uncaughtException', function(err) {
    console.log('Unhandled Exception', err);
});

process.on('exit', function(code) {
  console.log('About to exit with code:', code);
});