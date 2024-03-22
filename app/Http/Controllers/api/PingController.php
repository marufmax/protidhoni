<?php

namespace App\Http\Controllers\api;

use App\Http\Controllers\Controller;

class PingController
{
    public function __invoke()
    {
        return [
          'ping' => 'pong',
        ];
    }
}
