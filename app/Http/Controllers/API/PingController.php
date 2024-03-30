<?php

namespace App\Http\Controllers\API;

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
