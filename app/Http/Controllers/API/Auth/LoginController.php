<?php

namespace App\Http\Controllers\API\Auth;

use App\Http\Controllers\Controller;
use App\Http\Requests\LoginRequest;
use App\Http\Requests\RegisterRequest;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\Auth;

class LoginController extends Controller
{
    private const TOKEN_PREFIX = 'api_token';

    public function login(LoginRequest $request): JsonResponse
    {
        if (!Auth::once($request->only(['email', 'password']))) {
            return response()->json([
                'error' => 'invalid user',
            ], Response::HTTP_UNAUTHORIZED);
        }

        $user = User::where(['email' => $request->get('email')])->firstOrFail();

        return response()->json([
            'user' => [
                'id' => $user->id,
                'name' => $user->name,
                'email' => $user->email,
            ],
            'token' => $user
                ->createToken(sprintf('%s_of_%s', self::TOKEN_PREFIX, $user->email))
                ->plainTextToken,
        ], Response::HTTP_CREATED);
    }

    public function register(RegisterRequest $request): JsonResponse
    {
        $user = User::create($request->only(['email', 'password', 'name']));

        if (!Auth::once($request->only('email', 'password'))) {
            return response()->json([
               'error' => 'user registered but failed to authenticate',
            ], Response::HTTP_FAILED_DEPENDENCY);
        }

        return response()->json([
           'user' => [
               'id' => $user->id,
               'name' => $user->name,
               'email' => $user->email,
           ],
           'token' => $user
               ->createToken(sprintf('%s_of_%s', self::TOKEN_PREFIX, $user->email))
               ->plainTextToken,
        ]);
    }
}
