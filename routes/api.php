<?php

use App\Http\Controllers\API\Auth\LoginController;
use App\Http\Controllers\API\Auth\PingController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', static function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');

Route::get('ping', PingController::class);

## Auth
Route::post('login', [LoginController::class, 'login']);
Route::post('register', [LoginController::class, 'register']);

## Chat

