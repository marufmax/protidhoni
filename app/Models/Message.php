<?php

namespace App\Models;

use App\Enums\MessageStatus;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Support\Facades\Auth;

class Message extends Model
{
    protected $fillable = [
        'contents',
        'status',
        'read_at',
        'sent_to',
    ];

    protected $attributes = [
      'status' => MessageStatus::Pending
    ];

    protected $casts = [
        'status' => MessageStatus::class
    ];

    public static function boot()
    {
        parent::boot();

        static::creating(static fn ($message) => $message->sent_from = Auth::id());
    }

    public function sender(): BelongsTo
    {
        return $this->belongsTo(User::class, 'sent_from');
    }

    public function receiver(): BelongsTo
    {
        return $this->belongsTo(User::class, 'sent_to');
    }
}
