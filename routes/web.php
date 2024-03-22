<?php

use App\Http\Controllers\api\PingController;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/api/ping', PingController::class);
