<?php

namespace App\Enums;

enum MessageStatus: string
{
    case Pending = 'pending';
    case Sent = 'sent';
    case Received = 'received';
    case Failed = 'failed';
    case Read = 'read';
}
