<?php

use App\Http\Controllers\API\PingController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', static function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');

Route::get('ping', PingController::class);
Route::post('login', [App\Http\Controllers\API\LoginController::class, 'create']);
